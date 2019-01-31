def _impl(ctx):
  script = "%s -f %s -m %s -I /,.,vendor/github.com/googleapis/googleapis %s" % (
    ctx.files._pbvalidate[0].short_path,
    ctx.files.proto[0].short_path,
    ctx.attr.message,
    ctx.files.src[0].short_path,
  )

  # Write the file, it is executed by 'bazel test'.
  ctx.actions.write(
      output = ctx.outputs.executable,
      content = script,
  )

  # To ensure the files needed by the script are available, we put them in
  # the runfiles.
  runfiles = ctx.runfiles(files = ctx.files.src 
                                + ctx.files.proto
                                + ctx.files.deps
                                + ctx.files._googleapi
                                + ctx.files._pbvalidate)
  return [DefaultInfo(runfiles = runfiles)]

pbvalidate_test = rule(
  implementation = _impl,
  test = True,
  attrs = {
    "src": attr.label(allow_files = True),
    "proto": attr.label(allow_files = True),
    "deps": attr.label_list(allow_files = True),
    "message": attr.string(),

    # special case for vendored googleapis library; added to the search path
    "_googleapi": attr.label(default = Label("//vendor/github.com/googleapis/googleapis:protos")),

    # validation tool
    "_pbvalidate": attr.label(
      default = Label("//dev/pbvalidate"),
      executable = True,
      cfg = "host",
    ),
  }
)
