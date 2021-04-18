package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const POWER_SHELL = `$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
     param($commandName, $wordToComplete, $cursorPosition)
     $other = "$wordToComplete --generate-bash-completion"
         Invoke-Expression $other | ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
         }
 }`

func psProfileGen(binary string) string {
	return "\r\n& $(\"$(Split-Path -Path $profile)/AutoComplete/" + binary + ".ps1\")\r\n"
}

func setupPowerShell(profile string) error {
	// Create AutoComplete Folder
	dir := filepath.Dir(profile)
	folder := path.Join(dir, "AutoComplete")
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return err
	}

	for _, v := range BINARY_NAME {
		err = write(path.Join(folder, v+".ps1"), POWER_SHELL, false)
		if err != nil {
			return err
		}
	}
	err = clearExisting(profile)
	if err != nil {
		return err
	}
	for _, v := range BINARY_NAME {
		err = appendTo(profile, psProfileGen(v), false)
		if err != nil {
			return err
		}
	}
	fmt.Println("Please either restart your shell or run:")
	fmt.Println("\t& $profile")
	return nil
}

func clearExisting(profile string) error {
	for _, v := range BINARY_NAME {
		err := removeFromFile(profile, psProfileGen(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func tearDownPowerShell(profile string) error {
	dir := filepath.Dir(profile)
	folder := path.Join(dir, "AutoComplete")
	for _, v := range BINARY_NAME {
		err := os.Remove(path.Join(folder, v+".ps1"))
		if err != nil {
			fmt.Println(err)
		}
	}

	return clearExisting(profile)
}
