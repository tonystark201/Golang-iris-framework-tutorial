/*
 * @Descripttion: Do not edit
 * @version: v0.1.0
 * @Author: TSZ201
 * @Date: 2021-02-27 23:17:19
 * @LastEditors: TSZ201
 * @LastEditTime: 2021-02-27 23:17:20
 */
package commons

const TokenPrefix string = "ABCD"
const TokenDelimiter string = "&&"

type schemaPath struct {
	TeacherSchemaPath string
}

var SchemaPath = &schemaPath{}

func init() {
	SchemaPath.TeacherSchemaPath = "file:///iris_demo/schema/teacher.json"
}
