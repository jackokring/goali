# Welcome to Random Markdown

Is it really connected to [The Wiki](https://github.com/jackokring/goali/wiki)? Who knows? Is it [The Blog](https://jackokring.github.io/goali/)? I don't think it is. I suppose the blog is like a promotional kind of page with styling and a bit of keyword indexing. The wiki is more informal and is supposed to be like a wikipedia of the whole github repo. Other people can add to the wiki sometimes if it is so enabled.

Strangely, the wiki is not part of the repo download pull, but the blog is. So that makes Random Markdown like an offline set of textual information, without template styling or wiki searching. The wiki can't be edited offline unlike the blog and Random Markdown can. It's a magical information triad. So what am I going to be putting here? Timeless classics? Ah, we shall see.

## `xdg-desktop-portal` Both the `-gtk` and the `-xapp`

Ah, the XApp absractions and the `gtk3-nocsd` debacle. I'm not sure why you'd move to Gtk4. There's even `cros-xdg-desktop-portal` for those who've done a search. Seems strange to name it postfix on the system. A classic in commutative product naming file/location systems which hasn't happened circa 2024. I'm not a fan of CSD things, but not for asthetic reasons. I just never want an app to be able to escape its clip box and draw some pretend widgets to fish for data. Such a thing as CSD should be optional, and enforcable SSD by a desktop environment setting. You wouldn't want to download a full screen game, exit, and press a "hardwired" key shortcut to disable all CSD, just to be sure it was the right close box, and not an AI recalled sucker trojan on the end of the game, would you?

## I'm Enjoying Using Emacs, Just Like I Used to

Since doing an expansion of CUA mode by an `init.el` with some more muscle memory key combinations, it's been much "easier" to use. I could almost disable the tool bar as I'm picking it up, and have `which-key` help installed. A little theme addition, and perhaps some language modes plus some other tooling, and I'm almost regretting not banishing VSCode earlier. It's not as fast as Nano, but sure is faster than VSCode. Also Nano has some operational slowness even though it loads very fast for quick edits on small files. If anything Nano could also do with somekind of CUA overhaul too. I guess there's always the `mc` editor too, but I just don't like `vi(m)` which sure can be customized, but it has such an unfamiliar key layout steming back to "very old" line editors. Emacs and its embedded LISP is positively 1960s in comparison, with lots of 80s and 90s additions.

## How to Export a Theme, Maybe or Already

Wouldn't it be nice just to have a few command line switches for any app say `--supports-exporting-theme <feature>` and `--export-theme-xdg <feature>` such that one checks if it would work by returning an exit code other than zero, and the other actually alters the theme using the `xdg-desktop-portal`. The hard part would be making sure bad command switches are reported as errors in all tools. This brings me back to Emacs. It accepts, the bad option, as it has an internal LISP and so has the option to attach handlers onto command line options. So technically it would need to have the option to exit on a bad command line option. Apparently there is `command-switch-alist` which should have dotted pairs `("--option" . function-to-call)` added to handle options. The `function-to-call` recieves `"--option"` as its argument, and `command-line-args-left` contains the remaing command line, which can be modified to consume the `cdr` if needed for a `<fearure>`. Then there's `(exit 0)` for handling a return code status with zero being true (an ironic situation of negative logic in the shell).

I mean it's not as though `icon-set` or `mouse-icons` would be able to come from themes that are colour based. Some themes would just be a colour list to match, while others would be some weird CSS. The mismatched CSD, just becoming an option to theme might be more problematic given the security need to cancel decorations exceeding the clip box. Notice how browser information dialogs do not centre, but obscure some of your toolbar screen estate? Ah the pop-up hell of the 90s incarnate.

## Writing Time Independnt Code

Languages. They are somewhat interesting. Some of the features include prepositions as auxilliary verbs, word order designed to illicit imagination, or dumb found a lingo of fish rife elle ah re add. How se aid, kinda no vailable insruct to be maid. It's a pizza, nun left? Sew, you could arse-k yourself? "Watt is to be do-né ear?", and wi' t' sigh lent 'ate shhhh.

They get built as sure as time follows desire. The wants flowing into the universal Bob of the infinite bank of build. Selling shorts, for finer weathers or knots.

Yours In-Seerly
*Simon "Jacko" Jackson, BEng.*
