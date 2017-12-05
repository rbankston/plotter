# plotter
Keeps track of where your kubeconfigs live

## Problem
A lot of tools rely on your kubeconfig and some people do not like using contexts other than for namespaces. This is true for people with multiple clusters and wanting to verify exactly what clusters they are currently accessing.

## plotter's solution
plotter's set up and use relies on a plotter folder located at `.kube/plotter` by default. Plotter has two files at a time, the active file with its alias and backup. When you switch plotter files plotter checks backup and the active file for changes and asks you which one you want to keep. 

## plotter installation
to be determined

## getting started

```
plotter init <name>
```

plotter init a name is how to get started. The init creates the `.kube/plotter` directory to hold your kubeconfigs. The name takes your current `.kube/config` file and moves it to `.kube/plotter/<name>` along with creating `.kube/plotter/backup` for the backup function.

## commands

```
plotter list
```
plotter list shows all of the plotter configs available to use with their name.


```
plotter use <name>
```
plotter use <name> makes that specific file the active plotter file. This also verifies if any changes were made to your previous plotter config and which one to save or to create a new plotter file out the back up leaving the previous file the same without modifications.

```
plotter
```
plotter gives you the name of the active plotter file.

```
plotter init <name>
```
plotter init is how you get started. This command checks if a plotter directory exists and if not creates the directory and The name takes your current `.kube/config` file and moves it to `.kube/plotter/<name>` along with creating `.kube/plotter/backup` for the backup function.