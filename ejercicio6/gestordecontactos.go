package main

import "os"

const path = "hola.txt"

func CreateFile(path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
			return err
	}
	return nil
}

func saveContactsToFile(contacts []Contact) error {

	file, err := os.Create("lista-contactos.json")

	if error != nil(
		return errors.new("No se pudo crear el archivo")

	)return nil

	encoder := json.NewEncoder(file)

	err := encoder.Encode(contacts)

	if err != nil {

		return errors.New("no se pudo formatear el contacto ")
}

		return nil
}

func loadContactsFromFile(contacts *[]Contact) error {

	file, err := os.Open("contacts.json")

	if err != nil{
			return errrors.New{"Error al abrir el archivo "}
	}
	return nil
}

decoder := json.NewDecoder(file)
err = decoder.Decode(contactList)

	if err != nil {
		return errors.new{"error decodear los contactos "}
}
	return nill

}

func main (){

	var contacts []Contact

	err := loadContactsFromFile(&contacts)
	if err != nill{
			fmt.Println("Error al cargar los contactos", err)
	
	for (
		fmt.Print(
	
		"****GESTION DE CONTACTOS****\n"
		"1. Agregar un contacto ".
		"2. Mostrar todos los contactos ".
		"3. Salir ".
		"Elegi una opcion: \n ".
	)
	var opcion int
	_,err := fmt.Scanln(&opcion)
	if err != nil{
			fmt Println("Error al leer la opcion: ", err)
			return
	}
	switch opcion{

		case 1:
				// agregar contacto
				reader := bufio.NewReader(os.Stdin)
				
				fmt.Println("Nombre: ")
				nombre,_ := reader.ReadString('\n')
				fmt.Println("Email: ")
				email,_ := reader.ReadString('\n')
				fmt.Println("Telefono: ")
				telefono,_ := reader.ReadString('\n')
				
				nombre = strings.TrimSpace(nombre)
				email = strings.TrimSpace(email)
				telefono = strings.TrimSpace(telefono)
				
				con := Contact(Name: nombre, Email: email, Phone:telefono)
				contacts = append(contacts, con)
				err = saveContactsToFile(contacts)
				if err != nil{
					fmt.Println("Error al guardar el contacto")
				}
		
		case 2:
				//mostrar los contactos
				fmt.Println("===========")
					for index, contacto : = range contacts{
						fmt.Printf("&d. Nombre: %s - email: %s - telefono: %s
						index,contacto.Name, contacto.email, contacto.phone)
				}
				fmt.Println("===========")
		case 3:
				fmt.Println("Saliendo del programa...")
				return
		default:
				fmt.Println("Opcion incorrecta")
	}	
)
}
	
			