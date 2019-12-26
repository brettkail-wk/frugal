/*
 * Copyright 2017 Workiva
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package globals

import (
	"fmt"
)

// Version of the Frugal compiler.
const Version = "3.4.8"

// PrintWarning prints the given message to stdout in yellow font.
func PrintWarning(msg string) {
	fmt.Println("\x1b[33m" + msg + "\x1b[0m")
}
