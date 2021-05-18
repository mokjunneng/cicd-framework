#!/bin/bash

# TODO: Enhance CLI arguments parsing, include help method
for i in "$@"
do

case $i in
    --build_folder=*)
    BUILD_FOLDER="${i#*=}"
    shift # past argument=value
    ;;
    --artifact_path=*)
    ARTIFACT_PATH="${i#*=}"
    shift # past argument=value
    ;;
    --package_name=*)
    PACKAGE_NAME="${i#*=}"
    shift # past argument=value
    ;;
    --copyright=*)
    COPYRIGHT="${i#*=}"
    shift # past argument=value
    ;;
    --email=*)
    EMAIL="${i#*=}"
    shift # past argument=value
    ;;
    --author=*)
    AUTHOR="${i#*=}"
    shift # past argument=value
    ;;
    --section=*)
    SECTION="${i#*=}"
    shift # past argument=value
    ;;
    --ci_project_url=*)
    CI_PROJECT_URL="${i#*=}"
    shift # past argument=value
    ;;
    --ci_repository_url=*)
    CI_REPOSITORY_URL="${i#*=}"
    shift # past argument=value
    ;;
    --short_description=*)
    SHORT_DESCRIPTION="${i#*=}"
    shift # past argument=value
    ;;
    --long_description=*)
    LONG_DESCRIPTION="${i#*=}"
    shift # past argument=value
    ;;
    --target_ppa=*)
    TARGET_PPA="${i#*=}"
    shift # past argument=value
    ;;
    *)
          # unknown option
    ;;
esac
done

# Copy all Linux binaries to new folder
mkdir -p ${BUILD_FOLDER}
cp ${ARTIFACT_PATH}/*Linux* ${BUILD_FOLDER}/
cd ${BUILD_FOLDER}
gunzip *.gz

# import gpg
gpg || true
gpg --import ${GPG_PRIVATE_KEY}
gpg --import ${GPG_PUBLIC_KEY}
gpg --import-ownertrust ${GPG_OWNERTRUST}

# Install required tools
apt install gnupg dput dh-make devscripts lintian

# dh make
dh_make -p ${PACKAGE_NAME} --single --native --copyright ${COPYRIGHT} --email ${EMAIL}
rm debian/*.ex debian/*.EX # these files are not needed

# update control and changelog
perl -i -pe "s/unstable/$(lsb_release -cs)/" debian/changelog
perl -i -pe 's/^(Section:).*/$1 ${SECTION}/' debian/control
perl -i -pe 's/^(Homepage:).*/$1 ${CI_PROJECT_URL}/' debian/control
perl -i -pe 's/^#(Vcs-Browser:).*/$1 ${CI_PROJECT_URL}/' debian/control
perl -i -pe 's/^#(Vcs-Git:).*/$1 ${CI_REPOSITORY_URL}/' debian/control
perl -i -pe 's/^(Description:).*/$1 ${SHORT_DESCRIPTION}/'
perl -i -pe $'s/^ <insert long description.*/ ${LONG_DESCRIPTION}/' debian/control
perl -i -pe 's/^(Standards-Version:) 3.9.6/$1 3.9.7/' debian/control
perl -i -0777 -pe "s/(Copyright: ).+\n +.+/\${1}$(date +%Y) ${AUTHOR} ${EMAIL}/" debian/copyright

# Build the package
debuild -S | tee /tmp/debuild.log 2>&1

# Upload the package
dput ppa:${TARGET_PPA} "$(perl -ne 'print $1 if /dpkg-genchanges -S >(.*)/' /tmp/debuild.log)"  	# uses log file from previous section