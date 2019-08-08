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
	// "github.com/eclipse/codewind-installer/utils"
	// "github.com/urfave/cli"
)

func ValidateCommand(projectPath string) string {
	fmt.Println(projectPath)
	var projectType string
	if _, err := os.Stat(path.Join(projectPath, "pom.xml")); err == nil {
		projectType = "java"
	}
	if _, err := os.Stat(path.Join(projectPath, "package.json")); err == nil {
		projectType = "nodejs"
	}
	if _, err := os.Stat(path.Join(projectPath, "Package.swift")); err == nil {
		projectType = "swift"
	}
	fmt.Println(projectType)
	return projectType
}
