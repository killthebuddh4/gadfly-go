#!/bin/sh

cd vscode

vsce package

code --install-extension *.vsix

cd ..