package file

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var RenameSimilarityThresholdChange = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Change the rename similarity threshold while in the files panel",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("original", "one\ntwo\nthree\nfour\nfive\n")
		shell.Commit("add original")

		shell.DeleteFileAndAdd("original")
		shell.CreateFileAndAdd("renamed", "one\ntwo\nthree\nfour\nfive\nsix\nseven\neight\nnine\nten\n")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsFocused().
			Lines(
				Equals("▼ /"),
				Equals("  D  original"),
				Equals("  A  renamed"),
			).
			Press(keys.Universal.DecreaseRenameSimilarityThreshold).
			Tap(func() {
				t.ExpectToast(Equals("Changed rename similarity threshold to 45%"))
			}).
			Lines(
				Equals("R  original → renamed"),
			).
			Press(keys.Universal.FocusMainView).
			Tap(func() {
				t.Views().Main().
					Press(keys.Universal.IncreaseRenameSimilarityThreshold)
				t.ExpectToast(Equals("Changed rename similarity threshold to 50%"))
			}).
			Lines(
				Equals("▼ /"),
				Equals("  D  original"),
				Equals("  A  renamed"),
			)
	},
})
