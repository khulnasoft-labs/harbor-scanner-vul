# That's the only place where you're supposed to specify version of Vul.
ARG VUL_VERSION=0.44.0

FROM khulnasoft/vul:${VUL_VERSION}

# An ARG declared before a FROM is outside of a build stage, so it can't be used in any
# instruction after a FROM. To use the default value of an ARG declared before the first
# FROM use an ARG instruction without a value inside of a build stage.
ARG VUL_VERSION

RUN adduser -u 10000 -D -g '' scanner scanner

COPY scanner-vul /home/scanner/bin/scanner-vul

ENV VUL_VERSION=${VUL_VERSION}

USER scanner

ENTRYPOINT ["/home/scanner/bin/scanner-vul"]
