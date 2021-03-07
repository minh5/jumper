# jumper

CLI tool that *jumps* to directory based on a `JUMP_TO` env var. Just a little project to learn Go. 

UPDATE: Found out that you can't just `cd` from Go because Go launches a child process to change directory and returns back to the parent. So I had to do a workaround inspired by [leap](https://github.com/Xercoy/leap), which involves having a bash function that calls it. Basically Go will just validate the env var.

Yes I am aware I can just do this 

```
function jumper {
   if [[ -z "$JUMP_TO" ]]; then
    echo "Must provide JUMP_TO" 1>&2
    exit 1
   fi

   cd $JUMP_TO
}
```

But it wouldn't be as fun.
