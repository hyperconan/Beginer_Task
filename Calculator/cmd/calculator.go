package main

import (
	"Calculator/internal"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name: "Calculator",
		//Flags: []cli.Flag{
		//	&cli.StringFlag{
		//		Name:  "a",
		//		Usage: "First number",
		//	},
		//},  // go calculator --a=10
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var expression string
			if cmd.NArg() > 0 {
				expression = cmd.Args().Get(0)
			}
			result, err := internal.DoExpression(expression)
			if err != nil {
				return err
			}
			fmt.Println(result)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
	}
}
