/**
 * Package prototype
 * @Author fengfeng.mei <Biophiliam@protonmail.com>
 * @Date 2024/12/29 14:33
 */

package main

// 以文件系统为例创建一个原型模式
func main() {
	file1 := &File{Name: "file1"}
	file2 := &File{Name: "file2"}
	file3 := &File{Name: "file3"}

	folder1 := &Folder{Name: "folder1", Children: []FileSystemInterface{file1, file2}}
	folder2 := &Folder{Name: "folder2", Children: []FileSystemInterface{folder1, file3}}
	folder3 := &Folder{Name: "folder3", Children: []FileSystemInterface{folder2}}

	cloneFolder3 := folder3.Clone()

	folder3.Print("   ")
	cloneFolder3.Print("   ")
}

// FileSystemInterface 创建一个文件系统
type (
	FileSystemInterface interface {
		Clone() FileSystemInterface // 克隆的文件系统
		Print(preFix string)        // 打印文件系统
	}
)

type (
	// Folder 文件夹
	Folder struct {
		Name     string
		Children []FileSystemInterface // 文件夹下还可以有文件夹
	}

	// File 文件
	File struct {
		Name string
	}
)

// Clone 当前文件夹
func (f *Folder) Clone() FileSystemInterface {
	cloneName := "_clone_" + f.Name
	var cloneChildren []FileSystemInterface

	if f.Children != nil {
		for _, child := range f.Children {
			cloneChildren = append(cloneChildren, child.Clone())
		}
	}

	return &Folder{
		Name:     cloneName,
		Children: cloneChildren,
	}
}

func (f *Folder) Print(preFix string) {
	println(preFix + f.Name)
	for _, child := range f.Children {
		child.Print(preFix + preFix)
	}
}

func (f *File) Clone() FileSystemInterface {
	cloneName := "_clone_" + f.Name
	return &File{
		Name: cloneName,
	}
}

func (f *File) Print(preFix string) {
	println(preFix + f.Name)
}
