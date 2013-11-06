package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var copyunpad_code cu.Function

type copyunpad_args struct {
	arg_dst unsafe.Pointer
	arg_Dx  int
	arg_Dy  int
	arg_Dz  int
	arg_src unsafe.Pointer
	arg_Sx  int
	arg_Sy  int
	arg_Sz  int
	argptr  [8]unsafe.Pointer
}

// Wrapper for copyunpad CUDA kernel, asynchronous.
func k_copyunpad_async(dst unsafe.Pointer, Dx int, Dy int, Dz int, src unsafe.Pointer, Sx int, Sy int, Sz int, cfg *config, str int) {
	if copyunpad_code == 0 {
		copyunpad_code = fatbinLoad(copyunpad_map, "copyunpad")
	}

	var _a_ copyunpad_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_Dx = Dx
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_Dx)
	_a_.arg_Dy = Dy
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Dy)
	_a_.arg_Dz = Dz
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_Dz)
	_a_.arg_src = src
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_Sx = Sx
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_Sx)
	_a_.arg_Sy = Sy
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_Sy)
	_a_.arg_Sz = Sz
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_Sz)

	args := _a_.argptr[:]
	cu.LaunchKernel(copyunpad_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)
}

// Wrapper for copyunpad CUDA kernel, synchronized.
func k_copyunpad(dst unsafe.Pointer, Dx int, Dy int, Dz int, src unsafe.Pointer, Sx int, Sy int, Sz int, cfg *config) {
	const stream = 0
	k_copyunpad_async(dst, Dx, Dy, Dz, src, Sx, Sy, Sz, cfg, stream)
	Sync(stream)
}

var copyunpad_map = map[int]string{0: "",
	20: copyunpad_ptx_20,
	30: copyunpad_ptx_30,
	35: copyunpad_ptx_35}

const (
	copyunpad_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry copyunpad(
	.param .u64 copyunpad_param_0,
	.param .u32 copyunpad_param_1,
	.param .u32 copyunpad_param_2,
	.param .u32 copyunpad_param_3,
	.param .u64 copyunpad_param_4,
	.param .u32 copyunpad_param_5,
	.param .u32 copyunpad_param_6,
	.param .u32 copyunpad_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<24>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copyunpad_param_0];
	ld.param.u32 	%r4, [copyunpad_param_1];
	ld.param.u32 	%r5, [copyunpad_param_2];
	ld.param.u32 	%r8, [copyunpad_param_3];
	ld.param.u64 	%rd4, [copyunpad_param_4];
	ld.param.u32 	%r6, [copyunpad_param_5];
	ld.param.u32 	%r7, [copyunpad_param_6];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 8 1
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 2 9 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 2 10 1
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 2 12 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 13 1
	mad.lo.s32 	%r18, %r3, %r7, %r2;
	mad.lo.s32 	%r19, %r18, %r6, %r1;
	mul.wide.s32 	%rd5, %r19, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mad.lo.s32 	%r20, %r3, %r5, %r2;
	mad.lo.s32 	%r21, %r20, %r4, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;

BB0_2:
	.loc 2 15 2
	ret;
}


`
	copyunpad_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry copyunpad(
	.param .u64 copyunpad_param_0,
	.param .u32 copyunpad_param_1,
	.param .u32 copyunpad_param_2,
	.param .u32 copyunpad_param_3,
	.param .u64 copyunpad_param_4,
	.param .u32 copyunpad_param_5,
	.param .u32 copyunpad_param_6,
	.param .u32 copyunpad_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<24>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copyunpad_param_0];
	ld.param.u32 	%r4, [copyunpad_param_1];
	ld.param.u32 	%r5, [copyunpad_param_2];
	ld.param.u32 	%r8, [copyunpad_param_3];
	ld.param.u64 	%rd4, [copyunpad_param_4];
	ld.param.u32 	%r6, [copyunpad_param_5];
	ld.param.u32 	%r7, [copyunpad_param_6];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 8 1
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 2 9 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 2 10 1
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 2 12 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 13 1
	mad.lo.s32 	%r18, %r3, %r7, %r2;
	mad.lo.s32 	%r19, %r18, %r6, %r1;
	mul.wide.s32 	%rd5, %r19, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mad.lo.s32 	%r20, %r3, %r5, %r2;
	mad.lo.s32 	%r21, %r20, %r4, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;

BB0_2:
	.loc 2 15 2
	ret;
}


`
	copyunpad_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry copyunpad(
	.param .u64 copyunpad_param_0,
	.param .u32 copyunpad_param_1,
	.param .u32 copyunpad_param_2,
	.param .u32 copyunpad_param_3,
	.param .u64 copyunpad_param_4,
	.param .u32 copyunpad_param_5,
	.param .u32 copyunpad_param_6,
	.param .u32 copyunpad_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<23>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copyunpad_param_0];
	ld.param.u32 	%r4, [copyunpad_param_1];
	ld.param.u32 	%r5, [copyunpad_param_2];
	ld.param.u32 	%r8, [copyunpad_param_3];
	ld.param.u64 	%rd4, [copyunpad_param_4];
	ld.param.u32 	%r6, [copyunpad_param_5];
	ld.param.u32 	%r7, [copyunpad_param_6];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 3 8 1
	mov.u32 	%r9, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	.loc 3 9 1
	mov.u32 	%r12, %ntid.y;
	mov.u32 	%r13, %ctaid.y;
	mov.u32 	%r14, %tid.y;
	mad.lo.s32 	%r2, %r12, %r13, %r14;
	.loc 3 10 1
	mov.u32 	%r15, %ntid.z;
	mov.u32 	%r16, %ctaid.z;
	mov.u32 	%r17, %tid.z;
	mad.lo.s32 	%r3, %r15, %r16, %r17;
	.loc 3 12 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r8;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_2;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 13 1
	mad.lo.s32 	%r18, %r3, %r7, %r2;
	mad.lo.s32 	%r19, %r18, %r6, %r1;
	mul.wide.s32 	%rd5, %r19, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.nc.f32 	%f1, [%rd6];
	mad.lo.s32 	%r20, %r3, %r5, %r2;
	mad.lo.s32 	%r21, %r20, %r4, %r1;
	mul.wide.s32 	%rd7, %r21, 4;
	add.s64 	%rd8, %rd1, %rd7;
	st.global.f32 	[%rd8], %f1;

BB2_2:
	.loc 3 15 2
	ret;
}


`
)
