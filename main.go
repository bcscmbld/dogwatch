package main

import (
	"fmt"
	"os"

	"github.com/gugahoi/dogwatch/pkg/cmd"
	"github.com/gugahoi/dogwatch/pkg/subcmd"
)

var version = "SNAPSHOT"

func main() {
	e := subcmd.Parse(
		usage,
		new(cmd.List),
	)
	os.Exit(e)
}

//nolint
func usage() {
	fmt.Printf(`
		:::::::::   ::::::::   ::::::::  :::       :::     ::: ::::::::::: ::::::::  :::    :::
		:+:    :+: :+:    :+: :+:    :+: :+:       :+:   :+: :+:   :+:    :+:    :+: :+:    :+:
		+:+    +:+ +:+    +:+ +:+        +:+       +:+  +:+   +:+  +:+    +:+        +:+    +:+
		+#+    +:+ +#+    +:+ :#:        +#+  +:+  +#+ +#++:++#++: +#+    +#+        +#++:++#++
		+#+    +#+ +#+    +#+ +#+   +#+# +#+ +#+#+ +#+ +#+     +#+ +#+    +#+        +#+    +#+
		#+#    #+# #+#    #+# #+#    #+#  #+#+# #+#+#  #+#     #+# #+#    #+#    #+# #+#    #+#
		#########   ########   ########    ###   ###   ###     ### ###     ########  ###    ###
	
	VERSION: %s
	`, version)
}
