package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gowarden/zulu"
)

var (
	completions      = []string{"bear\tan animal", "bearpaw\ta dessert", "dog", "unicorn\tmythical"}
	specialCharComps = []string{"at@", "equal=", "slash/", "colon:", "period.", "comma,", "letter"}
)

func getCompsFilteredByPrefix(prefix string) []string {
	var finalComps []string
	for _, comp := range completions {
		if strings.HasPrefix(comp, prefix) {
			finalComps = append(finalComps, comp)
		}
	}
	return finalComps
}

var rootCmd = &zulu.Command{
	Use: "testprog",
	Run: func(cmd *zulu.Command, args []string) {
		fmt.Println("rootCmd called")
	},
}

// ======================================================
// Set of commands that filter on the 'toComplete' prefix
// ======================================================
var prefixCmd = &zulu.Command{
	Use:   "prefix",
	Short: "completions filtered on prefix",
}

var defaultCmdPrefix = &zulu.Command{
	Use:   "default",
	Short: "Directive: default",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return getCompsFilteredByPrefix(toComplete), zulu.ShellCompDirectiveDefault
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noSpaceCmdPrefix = &zulu.Command{
	Use:   "nospace",
	Short: "Directive: no space",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return getCompsFilteredByPrefix(toComplete), zulu.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noSpaceCharCmdPrefix = &zulu.Command{
	Use:   "nospacechar",
	Short: "Directive: no space, with comp ending with special char @=/:.,",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		var finalComps []string
		for _, comp := range specialCharComps {
			if strings.HasPrefix(comp, toComplete) {
				finalComps = append(finalComps, comp)
			}
		}
		return finalComps, zulu.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noFileCmdPrefix = &zulu.Command{
	Use:   "nofile",
	Short: "Directive: nofilecomp",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return getCompsFilteredByPrefix(toComplete), zulu.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noFileNoSpaceCmdPrefix = &zulu.Command{
	Use:   "nofilenospace",
	Short: "Directive: nospace and nofilecomp",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return getCompsFilteredByPrefix(toComplete), zulu.ShellCompDirectiveNoFileComp | zulu.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Set of commands that do not filter on prefix
// ======================================================
var noPrefixCmd = &zulu.Command{
	Use:   "noprefix",
	Short: "completions NOT filtered on prefix",
}

var noSpaceCmdNoPrefix = &zulu.Command{
	Use:   "nospace",
	Short: "Directive: no space",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return completions, zulu.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noFileCmdNoPrefix = &zulu.Command{
	Use:   "nofile",
	Short: "Directive: nofilecomp",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return completions, zulu.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var noFileNoSpaceCmdNoPrefix = &zulu.Command{
	Use:   "nofilenospace",
	Short: "Directive: nospace and nofilecomp",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return completions, zulu.ShellCompDirectiveNoFileComp | zulu.ShellCompDirectiveNoSpace
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

var defaultCmdNoPrefix = &zulu.Command{
	Use:   "default",
	Short: "Directive: default",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return completions, zulu.ShellCompDirectiveDefault
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Command that completes on file extension
// ======================================================
var fileExtCmdPrefix = &zulu.Command{
	Use:   "fileext",
	Short: "Directive: fileext",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return []string{"yaml", "json"}, zulu.ShellCompDirectiveFilterFileExt
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Command that completes on the directories within the current directory
// ======================================================
var dirCmd = &zulu.Command{
	Use:   "dir",
	Short: "Directive: subdir",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return nil, zulu.ShellCompDirectiveFilterDirs
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Command that completes on the directories within the 'dir' directory
// ======================================================
var subDirCmd = &zulu.Command{
	Use:   "subdir",
	Short: "Directive: subdir",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return []string{"dir"}, zulu.ShellCompDirectiveFilterDirs
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Command that returns an error on completion
// ======================================================
var errorCmd = &zulu.Command{
	Use:   "error",
	Short: "Directive: error",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return completions, zulu.ShellCompDirectiveError
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

// ======================================================
// Command that wants an argument starting with a --
// Such an argument is possible following a '--'
// ======================================================
var dashArgCmd = &zulu.Command{
	Use:   "dasharg",
	Short: "Wants argument --arg",
	ValidArgsFunction: func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return []string{"--arg\tan arg starting with dashes"}, zulu.ShellCompDirectiveDefault
	},
	Run: func(cmd *zulu.Command, args []string) {},
}

func setFlags() {
	rootCmd.Flags().String("customComp", "", "test custom comp for flags")
	rootCmd.RegisterFlagCompletionFunc("customComp", func(cmd *zulu.Command, args []string, toComplete string) ([]string, zulu.ShellCompDirective) {
		return []string{"firstComp\tthe first value", "secondComp\tthe second value", "forthComp"}, zulu.ShellCompDirectiveNoFileComp
	})

	rootCmd.Flags().String("theme", "", "theme to use (located in /dir/THEMENAME/)")
	rootCmd.Flags().SetAnnotation("theme", zulu.BashCompSubdirsInDir, []string{"dir"})

	dashArgCmd.Flags().Bool("flag", false, "a flag")
}

func main() {
	rootCmd.AddCommand(newCompletionCmd(os.Stdout))
	setFlags()

	rootCmd.AddCommand(
		prefixCmd,
		noPrefixCmd,
		fileExtCmdPrefix,
		dirCmd,
		subDirCmd,
		errorCmd,
		dashArgCmd,
	)

	prefixCmd.AddCommand(
		noSpaceCmdPrefix,
		noSpaceCharCmdPrefix,
		noFileCmdPrefix,
		noFileNoSpaceCmdPrefix,
		defaultCmdPrefix,
	)

	noPrefixCmd.AddCommand(
		noSpaceCmdNoPrefix,
		noFileCmdNoPrefix,
		noFileNoSpaceCmdNoPrefix,
		defaultCmdNoPrefix,
	)

	rootCmd.Execute()
}
