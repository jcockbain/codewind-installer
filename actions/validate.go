
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
	"os"
	"path"
	"io/ioutil"
	"strings"
)

//ValidateCommand - Checks for build files in the project to determine language and buildType
func ValidateCommand(projectPath string) string {
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
	return projectType
}

func determineProjectFramework(projectPath string) string {
	pathToPomXML := path.Join(projectPath, "pom.xml")
	pomXMLContents, _err := ioutil.ReadFile(pathToPomXML)
	if _err !=  nil {
		return "docker"
	}
	pomXMLString := string(pomXMLContents)
	if strings.Contains(pomXMLString, "<groupId>org.springframework.boot</groupId>"){
		return "spring"
	}
	if strings.Contains(pomXMLString, "<groupId>org.eclipse.microprofile</groupId>"){
		return "liberty"
	}
	return "docker"
}
