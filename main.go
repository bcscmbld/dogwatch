package main

import (
	"fmt"
	"os"
	"time"

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
	BUILD DATE: %s
	`, version, time.Now().Format("2006-01-02"))
}
