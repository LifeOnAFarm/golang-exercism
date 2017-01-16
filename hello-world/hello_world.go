/* Name:		Seamus
	Date:		12/01/17
	Program:	Simple Hello World program
 */
package greeting




const testVersion = 3


func HelloWorld(name string) string {

	if name != ""{
		return "Hello, " + name + "!"
	} else{
		return "Hello, World!"
	}
}
