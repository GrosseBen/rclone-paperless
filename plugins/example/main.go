package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rclone/rclone/fs/config/configstruct"
	"github.com/rclone/rclone/fs/hash"
	"github.com/rclone/rclone/lib/pacer"
)

// Register with Fs
func init() {
	fs.Register(&fs.RegInfo{
		Name:        "example",
		Description: "Example plugin",
		NewFs:       NewFs,
		Options: []fs.Option{{
			Name: "example_option",
			Help: "An example option for the plugin",
			Default:  "default_value",
		}},
	})
}

// Fs represents a remote file system
type Fs struct {
	name     string
	root     string
	opt      Options
	features *fs.Features
	pacer    *pacer.Pacer
}

// Options defines the configuration for this plugin
type Options struct {
	ExampleOption string `config:"example_option"`
}

// NewFs constructs an Fs from the path and configuration
func NewFs(ctx context.Context, name, root string, m configmap.Mapper) (fs.Fs, error) {
	// Parse config into Options struct
	opt := new(Options)
	err := configstruct.Set(m, opt)
	if err != nil {
		return nil, err
	}

	f := &Fs{
		name:  name,
		root:  root,
		opt:   *opt,
		pacer: pacer.New(),
	}
	f.features = (&fs.Features{}).Fill(ctx, f)
	f.features.DuplicateFiles = false
	f.features.ReadMimeType = false
	f.features.WriteMimeType = false

	return f, nil
}

// Name of the remote (as passed into NewFs)
func (f *Fs) Name() string {
	return f.name
}

// Root of the remote (as passed into NewFs)
func (f *Fs) Root() string {
	return f.root
}

// String converts Fs to a string
func (f *Fs) String() string {
	return fmt.Sprintf("Example plugin: %s", f.root)
}

// Features returns the optional features of this Fs
func (f *Fs) Features() *fs.Features {
	return f.features
}

// Precision of the remote
func (f *Fs) Precision() time.Duration {
	return time.Second
}

// NewObject finds the Object at remote path
func (f *Fs) NewObject(ctx context.Context, remote string) (fs.Object, error) {
	return nil, fs.ErrorObjectNotFound
}

// List the objects and directories in dir into entries
func (f *Fs) List(ctx context.Context, dir string) (entries fs.DirEntries, err error) {
	return nil, fs.ErrorDirNotFound
}

// Put the object
func (f *Fs) Put(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) (fs.Object, error) {
	return nil, fmt.Errorf("Put not supported")
}

// Mkdir creates the directory if it doesn't exist
func (f *Fs) Mkdir(ctx context.Context, dir string) error {
	return fmt.Errorf("Mkdir not supported")
}

// Rmdir removes the directory
func (f *Fs) Rmdir(ctx context.Context, dir string) error {
	return fmt.Errorf("Rmdir not supported")
}

// Purge deletes all the files and directories
func (f *Fs) Purge(ctx context.Context, dir string) error {
	return fmt.Errorf("Purge not supported")
}

// Move src to this remote at dst
func (f *Fs) Move(ctx context.Context, src fs.Object, remote string) (fs.Object, error) {
	return nil, fmt.Errorf("Move not supported")
}

// DirMove moves src to this remote at dst
func (f *Fs) DirMove(ctx context.Context, src fs.Fs, remote string) error {
	return fmt.Errorf("DirMove not supported")
}

// Copy src to this remote at dst
func (f *Fs) Copy(ctx context.Context, src fs.Object, remote string) (fs.Object, error) {
	return nil, fmt.Errorf("Copy not supported")
}

// DirCopy copies src to this remote at dst
func (f *Fs) DirCopy(ctx context.Context, src fs.Fs, remote string) error {
	return fmt.Errorf("DirCopy not supported")
}

// About gets quota information
func (f *Fs) About(ctx context.Context) (*fs.Usage, error) {
	return nil, fmt.Errorf("About not supported")
}

// Hashes returns the supported hash sets
func (f *Fs) Hashes() hash.Set {
	return hash.NewHashSet(hash.None)
}

// CleanUp the Fs
func (f *Fs) CleanUp() {
}

// Shutdown the Fs
func (f *Fs) Shutdown() error {
	return nil
}

