package prompt

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"pkg.mattglei.ch/school/internal/conf"
)

type Answers struct {
	DocumentName string
	Class        conf.Class
	TemplatePath string
	Destination  string
}

func Ask(classes []conf.Class) (Answers, error) {
	var answers Answers
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().Title("Document Name").Value(&answers.DocumentName).WithHeight(1),
			huh.NewSelect[conf.Class]().Title("Class").OptionsFunc(func() []huh.Option[conf.Class] {
				opts := []huh.Option[conf.Class]{}
				for _, class := range classes {
					opts = append(
						opts,
						huh.NewOption(fmt.Sprintf("%s [%s]", *class.Code, *class.Name), class),
					)
				}
				return opts
			}, &answers.Class).WithHeight(len(classes)+2),
			huh.NewFilePicker().
				Title("LaTeX Template").
				CurrentDirectory("./templates/").
				ShowPermissions(false).
				ShowSize(true).
				DirAllowed(false).
				FileAllowed(true).Value(&answers.TemplatePath).
				Picking(true),
			huh.NewFilePicker().
				Title("Destination Folder").
				CurrentDirectory(answers.Class.Folder).
				DirAllowed(true).
				FileAllowed(false).
				Picking(true).
				Value(&answers.Destination),
		),
	).WithTheme(huh.ThemeBase())

	err := form.Run()
	if err != nil {
		return Answers{}, fmt.Errorf("%v failed to prompt user for form", err)
	}

	return answers, nil
}
