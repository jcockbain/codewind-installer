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

package utils

import (
	"encoding/json"
	"fmt"
)

// RemoveDuplicateEntries elements
func RemoveDuplicateEntries(inputArr []string) []string {

	encounteredElement := map[string]bool{}
	result := []string{}

	// Populate map if element != ""
	for _, element := range inputArr {
		if element != "" {
			encounteredElement[element] = true
		}
	}

	// Convert map => slice
	for key := range encounteredElement {
		result = append(result, key)
	}

	return result
}

func PrettyPrintJSON(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

// StringInSlice returns true if the string is in the slice
func StringInSlice(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}
