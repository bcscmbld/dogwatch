package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gugahoi/dogwatch/cmd"
)

var version = "SNAPSHOT"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err) // nolint: gas
		os.Exit(1)
	}
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
