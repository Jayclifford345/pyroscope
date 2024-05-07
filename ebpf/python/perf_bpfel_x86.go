// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package python

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type PerfErrorStats struct{ Errors [19]uint32 }

type PerfGlobalConfigT struct {
	BpfLogErr   uint8
	BpfLogDebug uint8
	Typecheck   uint8
	_           [5]byte
	NsPidIno    uint64
}

type PerfLibc struct {
	Musl                    bool
	_                       [1]byte
	PthreadSize             int16
	PthreadSpecific1stblock int16
}

type PerfPyEvent struct {
	K        PerfSampleKey
	StackLen uint32
	Stack    [96]uint32
	_        [4]byte
}

type PerfPyOffsetConfig struct {
	PyThreadStateFrame            int16
	PyThreadStateCframe           int16
	PyCFrameCurrentFrame          int16
	PyCodeObjectCoFilename        int16
	PyCodeObjectCoName            int16
	PyCodeObjectCoVarnames        int16
	PyCodeObjectCoLocalsplusnames int16
	PyCodeObjectCoNlocals         int16
	PyTupleObjectObItem           int16
	PyVarObjectObSize             int16
	PyObjectObType                int16
	PyTypeObjectTpName            int16
	VFrameCode                    int16
	VFramePrevious                int16
	VFrameLocalsplus              int16
	PyInterpreterFrameOwner       int16
	PyASCIIObjectSize             int16
	PyCompactUnicodeObjectSize    int16
	_                             [4]byte
	Base                          uint64
}

type PerfPyPidData struct {
	Offsets   PerfPyOffsetConfig
	Typecheck PerfPyTypecheckData
	Version   struct {
		Major uint32
		Minor uint32
		Patch uint32
	}
	Libc          PerfLibc
	_             [2]byte
	TssKey        int32
	CollectKernel uint8
	_             [7]byte
}

type PerfPySampleStateT struct {
	SymbolCounter int64
	Offsets       PerfPyOffsetConfig
	Typecheck     PerfPyTypecheckData
	Version       struct {
		Major uint32
		Minor uint32
		Patch uint32
	}
	CurCpu                 uint32
	FramePtr               uint64
	PythonStackProgCallCnt int64
	Sym                    PerfPySymbol
	Event                  PerfPyEvent
	Padding                uint64
}

type PerfPyStrType struct {
	Type           uint8
	SizeCodepoints uint8
}

type PerfPySymbol struct {
	Classname     [32]int8
	Name          [64]int8
	File          [128]int8
	ClassnameType PerfPyStrType
	NameType      PerfPyStrType
	FileType      PerfPyStrType
	Padding       PerfPyStrType
}

type PerfPyTypecheckData struct {
	PyCodeType                     int64
	PyFrameType                    int64
	PyBytesType                    int64
	PyUnicodeType                  int64
	PyTypeType                     int64
	PyDictType                     int64
	PyNoneType                     int64
	PyModuleType                   int64
	PyTupleType                    int64
	O_PyThreadStateDict            int64
	O_PyThreadStateInterp          int64
	SizePyThreadState              int64
	O_PyInterpreterStateTstateHead int64
	O_PyInterpreterStateFinalizing int64
	O_PyInterpreterStateModules    int64
	O_PyInterpreterStateImportlib  int64
	SizePyInterpreterState         int64
}

type PerfSampleKey struct {
	Pid       uint32
	Flags     uint32
	KernStack int64
	UserStack int64
}

// LoadPerf returns the embedded CollectionSpec for Perf.
func LoadPerf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PerfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Perf: %w", err)
	}

	return spec, err
}

// LoadPerfObjects loads Perf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*PerfObjects
//	*PerfPrograms
//	*PerfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadPerfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadPerf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// PerfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfSpecs struct {
	PerfProgramSpecs
	PerfMapSpecs
}

// PerfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfProgramSpecs struct {
	PyperfCollect   *ebpf.ProgramSpec `ebpf:"pyperf_collect"`
	ReadPythonStack *ebpf.ProgramSpec `ebpf:"read_python_stack"`
}

// PerfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type PerfMapSpecs struct {
	Counts       *ebpf.MapSpec `ebpf:"counts"`
	Events       *ebpf.MapSpec `ebpf:"events"`
	PyErrors     *ebpf.MapSpec `ebpf:"py_errors"`
	PyPidConfig  *ebpf.MapSpec `ebpf:"py_pid_config"`
	PyProgs      *ebpf.MapSpec `ebpf:"py_progs"`
	PyStateHeap  *ebpf.MapSpec `ebpf:"py_state_heap"`
	PySymbols    *ebpf.MapSpec `ebpf:"py_symbols"`
	PythonStacks *ebpf.MapSpec `ebpf:"python_stacks"`
	Stacks       *ebpf.MapSpec `ebpf:"stacks"`
}

// PerfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfObjects struct {
	PerfPrograms
	PerfMaps
}

func (o *PerfObjects) Close() error {
	return _PerfClose(
		&o.PerfPrograms,
		&o.PerfMaps,
	)
}

// PerfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfMaps struct {
	Counts       *ebpf.Map `ebpf:"counts"`
	Events       *ebpf.Map `ebpf:"events"`
	PyErrors     *ebpf.Map `ebpf:"py_errors"`
	PyPidConfig  *ebpf.Map `ebpf:"py_pid_config"`
	PyProgs      *ebpf.Map `ebpf:"py_progs"`
	PyStateHeap  *ebpf.Map `ebpf:"py_state_heap"`
	PySymbols    *ebpf.Map `ebpf:"py_symbols"`
	PythonStacks *ebpf.Map `ebpf:"python_stacks"`
	Stacks       *ebpf.Map `ebpf:"stacks"`
}

func (m *PerfMaps) Close() error {
	return _PerfClose(
		m.Counts,
		m.Events,
		m.PyErrors,
		m.PyPidConfig,
		m.PyProgs,
		m.PyStateHeap,
		m.PySymbols,
		m.PythonStacks,
		m.Stacks,
	)
}

// PerfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadPerfObjects or ebpf.CollectionSpec.LoadAndAssign.
type PerfPrograms struct {
	PyperfCollect   *ebpf.Program `ebpf:"pyperf_collect"`
	ReadPythonStack *ebpf.Program `ebpf:"read_python_stack"`
}

func (p *PerfPrograms) Close() error {
	return _PerfClose(
		p.PyperfCollect,
		p.ReadPythonStack,
	)
}

func _PerfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed perf_bpfel_x86.o
var _PerfBytes []byte
