/*******************************************************************************
 * Copyright (c) 2019 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/eclipse/codewind-installer/utils"
	"github.com/urfave/cli"
)

// DownloadTemplate using the url/link provided
func DownloadTemplate(c *cli.Context) {
	destination := c.Args().Get(0)

	if destination == "" {
		log.Fatal("destination not set")
	}

	repoURL := c.String("r")

	// expecting string in format 'https://github.com/<owner>/<repo>'
	if strings.HasPrefix(repoURL, "https://") {
		repoURL = strings.TrimPrefix(repoURL, "https://")
	}

	repoArray := strings.Split(repoURL, "/")
	owner := repoArray[1]
	repo := repoArray[2]
	branch := "master"

	var tempPath = ""
	const GOOS string = runtime.GOOS
	if GOOS == "windows" {
		tempPath = os.Getenv("TEMP") + "\\"
	} else {
		tempPath = "/tmp/"
	}

	zipURL := utils.GetZipURL(owner, repo, branch)

	time := time.Now().Format(time.RFC3339)
	time = strings.Replace(time, ":", "-", -1) // ":" is illegal char in windows
	tempName := tempPath + branch + "_" + time
	zipFileName := tempName + ".zip"

	// download files in zip format
	if err := utils.DownloadFile(zipFileName, zipURL); err != nil {
		log.Fatal(err)
	}

	// unzip into /tmp dir
	utils.UnZip(zipFileName, destination)

	//delete zip file
	utils.DeleteTempFile(zipFileName)
}

//ValidateProject type
func ValidateProject(c *cli.Context) {
	projectPath := c.Args().Get(0)
	language := "unknown"
	projectType := "docker"
	if _, err := os.Stat(path.Join(projectPath, "pom.xml")); err == nil {
		language = "java"
		projectType = determineProjectFramework(projectPath)
	}
	if _, err := os.Stat(path.Join(projectPath, "package.json")); err == nil {
		language = "nodejs"
		projectType = "nodejs"
	}
	if _, err := os.Stat(path.Join(projectPath, "Package.swift")); err == nil {
		language = "swift"
		projectType = "swift"
	}
	fmt.Println("project build type: " + projectType)
	fmt.Println("project language: " + language)
}

func determineProjectFramework(projectPath string) string {
	pathToPomXML := path.Join(projectPath, "pom.xml")
	pomXMLContents, _err := ioutil.ReadFile(pathToPomXML)
	if _err != nil {
		return "docker"
	}
	pomXMLString := string(pomXMLContents)
	if strings.Contains(pomXMLString, "<groupId>org.springframework.boot</groupId>") {
		return "spring"
	}
	if strings.Contains(pomXMLString, "<groupId>org.eclipse.microprofile</groupId>") {
		return "liberty"
	}
	return "docker"
}
