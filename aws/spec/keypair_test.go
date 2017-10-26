/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsspec

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func init() {
	tmpFolder, err := ioutil.TempDir("", "tmpfolder")
	if err != nil {
		panic(err)
	}
	genTestsCleanupFunc["createkeypair"] = func() { os.RemoveAll(tmpFolder) }
	keyGenerationFunction = func(size int, encryptKey bool) ([]byte, []byte, error) {
		var public string
		if encryptKey {
			public = "encrypted keypair"
		} else {
			public = "unencrypted keypair"
		}
		return []byte(public), []byte{}, nil
	}
	os.Setenv(keyDirEnv, tmpFolder)

	genTestsParams["createkeypair"] = map[string]interface{}{
		"name":      "my-kp",
		"encrypted": "true",
	}
	genTestsExpected["createkeypair"] = &ec2.ImportKeyPairInput{
		KeyName:           String("my-kp"),
		PublicKeyMaterial: []byte("encrypted keypair"),
	}

	genTestsOutputExtractFunc["createkeypair"] = func() interface{} {
		return &ec2.ImportKeyPairOutput{KeyName: String("my-kp")}
	}
	genTestsOutput["createkeypair"] = "my-kp"

	genTestsParams["deletekeypair"] = map[string]interface{}{
		"name": "my-kp",
	}
	genTestsExpected["deletekeypair"] = &ec2.DeleteKeyPairInput{
		KeyName: String("my-kp"),
	}
}

func (m *mockEc2) ImportKeyPair(input *ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	if got, want := input, genTestsExpected["createkeypair"]; !reflect.DeepEqual(got, want) {
		return nil, fmt.Errorf("got %#v, want %#v", got, want)
	}
	if outFunc, ok := genTestsOutputExtractFunc["createkeypair"]; ok {
		return outFunc().(*ec2.ImportKeyPairOutput), nil
	}
	return nil, nil
}
