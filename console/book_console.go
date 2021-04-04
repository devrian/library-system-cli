package console

import (
	"flag"
	"fmt"
	"library-system/helpers"
	"library-system/services"
	"os"
	"strconv"
)

type Console struct {
	bookService services.BookService
	commands    map[string]func() func(string) error
}

func NewBookConsole(bookService services.BookService) Console {
	c := Console{
		bookService: bookService,
	}

	c.commands = map[string]func() func(string) error{
		"add":    c.addBook,
		"rent":   c.rentBook,
		"return": c.returnBook,
		"rented": c.rentedBook,
		"list":   c.listBook,
		"get":    c.getBook,
	}

	return c
}

func (c Console) Console() error {
	cmdName := os.Args[1]
	cmd, ok := c.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command '%s'\n", cmdName)
	}

	return cmd()(cmdName)
}

func (c Console) Help() {
	var help string
	for name := range c.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("Usage of %s:\n<command> [<args>]\n%s", os.Args[0], help)
}

func (c Console) parseCommand(command *flag.FlagSet) error {
	err := command.Parse(os.Args[2:])
	if err != nil {
		return helpers.WrapError("could not parse '"+command.Name()+"' command flags", err)
	}

	return nil
}

func (c Console) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}
	if len(os.Args)-2 < minArgs {
		fmt.Printf("incorrect use of %s\n%s %s --help\n", os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d arg(s), %d provided", os.Args[1], minArgs, len(os.Args)-2)
	}

	return nil
}

func (c Console) reminderFlagsAddBook(f *flag.FlagSet) (string, string) {
	code, name := "", ""
	f.StringVar(&code, "code_of_book", "", "Reminder code")
	f.StringVar(&name, "name_of_book", "", "Reminder name")

	return code, name
}

func (c Console) reminderFlagsCode(f *flag.FlagSet) string {
	code := ""
	f.StringVar(&code, "code_of_book", "", "Reminder code")

	return code
}

func (c Console) addBook() func(string) error {
	return func(cmd string) error {
		addCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		c.reminderFlagsAddBook(addCmd)

		if err := c.checkArgs(2); err != nil {
			return err
		}
		if err := c.parseCommand(addCmd); err != nil {
			return err
		}

		idConvert, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return helpers.WrapError("could not add new book", err)
		}

		result, err := c.bookService.CreateBook(idConvert, os.Args[3], false)
		if err != nil {
			return helpers.WrapError("could not add new book", err)
		}

		fmt.Printf("add new book %s successfully\n", result.Name)
		return nil
	}
}

func (c Console) rentBook() func(string) error {
	return func(cmd string) error {
		rentCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		c.reminderFlagsCode(rentCmd)

		if err := c.checkArgs(1); err != nil {
			return err
		}
		if err := c.parseCommand(rentCmd); err != nil {
			return err
		}

		idConvert, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return helpers.WrapError("could not rent a book", err)
		}

		result, err := c.bookService.FlagBookRentByIDAndIsRent(idConvert, true)
		if err != nil {
			return helpers.WrapError("could not rent a book", err)
		}

		fmt.Printf("rent a book %s successfully\n", result.Name)
		return nil
	}
}

func (c Console) returnBook() func(string) error {
	return func(cmd string) error {
		returnCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		c.reminderFlagsCode(returnCmd)

		if err := c.checkArgs(1); err != nil {
			return err
		}
		if err := c.parseCommand(returnCmd); err != nil {
			return err
		}

		idConvert, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return helpers.WrapError("could not return a book", err)
		}

		result, err := c.bookService.FlagBookRentByIDAndIsRent(idConvert, false)
		if err != nil {
			return helpers.WrapError("could not return a book", err)
		}

		fmt.Printf("return a book %s successfully\n", result.Name)
		return nil
	}
}

func (c Console) rentedBook() func(string) error {
	return func(cmd string) error {
		listCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		if err := c.parseCommand(listCmd); err != nil {
			return err
		}

		books, err := c.bookService.GetAllBooksRented()
		if err != nil {
			return helpers.WrapError("could not list all book", err)
		}

		for _, book := range books {
			status := "return"
			if book.IsRented == true {
				status = "rented"
			}
			fmt.Println("--------------------------")
			fmt.Printf("book's code : %v\nbook's name : %s\nbook's status : %s\n", book.ID, book.Name, status)
		}

		return nil
	}
}

func (c Console) listBook() func(string) error {
	return func(cmd string) error {
		listCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		if err := c.parseCommand(listCmd); err != nil {
			return err
		}

		books, err := c.bookService.GetAll()
		if err != nil {
			return helpers.WrapError("could not list all book", err)
		}

		for _, book := range books {
			status := "return"
			if book.IsRented == true {
				status = "rented"
			}
			fmt.Println("--------------------------")
			fmt.Printf("book's code : %v\nbook's name : %s\nbook's status : %s\n", book.ID, book.Name, status)
		}

		return nil
	}
}

func (c Console) getBook() func(string) error {
	return func(cmd string) error {
		getCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		c.reminderFlagsCode(getCmd)

		if err := c.checkArgs(1); err != nil {
			return err
		}
		if err := c.parseCommand(getCmd); err != nil {
			return err
		}

		idConvert, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return helpers.WrapError("could not get a book", err)
		}

		result, err := c.bookService.GetBookByID(idConvert)
		if err != nil {
			return helpers.WrapError("could not get a book", err)
		}

		status := "return"
		if result.IsRented == true {
			status = "rented"
		}

		fmt.Println("detail of book")
		fmt.Println("-------------------------")
		fmt.Printf("book's code : %v\nbook's name : %s\nbook's status : %s\n", result.ID, result.Name, status)
		return nil
	}
}
