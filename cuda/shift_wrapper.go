package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var shift_code cu.Function

type shift_args struct {
	arg_dst unsafe.Pointer
	arg_src unsafe.Pointer
	arg_Nx  int
	arg_Ny  int
	arg_Nz  int
	arg_shx int
	arg_shy int
	arg_shz int
	argptr  [8]unsafe.Pointer
}

// Wrapper for shift CUDA kernel, asynchronous.
func k_shift_async(dst unsafe.Pointer, src unsafe.Pointer, Nx int, Ny int, Nz int, shx int, shy int, shz int, cfg *config, str int) {
	if shift_code == 0 {
		shift_code = fatbinLoad(shift_map, "shift")
	}

	var _a_ shift_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_src = src
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_Nx = Nx
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Nx)
	_a_.arg_Ny = Ny
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_Ny)
	_a_.arg_Nz = Nz
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_Nz)
	_a_.arg_shx = shx
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_shx)
	_a_.arg_shy = shy
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_shy)
	_a_.arg_shz = shz
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_shz)

	args := _a_.argptr[:]
	cu.LaunchKernel(shift_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)
}

// Wrapper for shift CUDA kernel, synchronized.
func k_shift(dst unsafe.Pointer, src unsafe.Pointer, Nx int, Ny int, Nz int, shx int, shy int, shz int, cfg *config) {
	const stream = 0
	k_shift_async(dst, src, Nx, Ny, Nz, shx, shy, shz, cfg, stream)
	Sync(stream)
}

var shift_map = map[int]string{0: "",
	20: shift_ptx_20,
	30: shift_ptx_30,
	35: shift_ptx_35}

const (
	shift_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<38>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r4, [shift_param_2];
	ld.param.u32 	%r5, [shift_param_3];
	ld.param.u32 	%r6, [shift_param_4];
	ld.param.u32 	%r7, [shift_param_5];
	ld.param.u32 	%r8, [shift_param_6];
	ld.param.u32 	%r9, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 9 1
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	.loc 2 10 1
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	.loc 2 11 1
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r6;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 14 1
	sub.s32 	%r19, %r3, %r9;
	mov.u32 	%r20, 0;
	.loc 3 238 5
	max.s32 	%r21, %r19, %r20;
	.loc 2 14 1
	add.s32 	%r22, %r6, -1;
	.loc 3 210 5
	min.s32 	%r23, %r21, %r22;
	.loc 2 14 1
	sub.s32 	%r24, %r2, %r8;
	.loc 3 238 5
	max.s32 	%r25, %r24, %r20;
	.loc 2 14 1
	add.s32 	%r26, %r5, -1;
	.loc 3 210 5
	min.s32 	%r27, %r25, %r26;
	.loc 2 14 1
	mad.lo.s32 	%r28, %r23, %r5, %r27;
	sub.s32 	%r29, %r1, %r7;
	.loc 3 238 5
	max.s32 	%r30, %r29, %r20;
	.loc 2 14 1
	add.s32 	%r31, %r4, -1;
	.loc 3 210 5
	min.s32 	%r32, %r30, %r31;
	.loc 2 14 1
	mad.lo.s32 	%r33, %r28, %r4, %r32;
	mul.wide.s32 	%rd5, %r33, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mad.lo.s32 	%r34, %r3, %r5, %r2;
	mad.lo.s32 	%r35, %r34, %r4, %r1;
	mul.wide.s32 	%rd7, %r35, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;

BB0_2:
	.loc 2 16 2
	ret;
}


`
	shift_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<38>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r4, [shift_param_2];
	ld.param.u32 	%r5, [shift_param_3];
	ld.param.u32 	%r6, [shift_param_4];
	ld.param.u32 	%r7, [shift_param_5];
	ld.param.u32 	%r8, [shift_param_6];
	ld.param.u32 	%r9, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 9 1
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	.loc 2 10 1
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	.loc 2 11 1
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r6;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_2;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 14 1
	sub.s32 	%r19, %r3, %r9;
	mov.u32 	%r20, 0;
	.loc 3 238 5
	max.s32 	%r21, %r19, %r20;
	.loc 2 14 1
	add.s32 	%r22, %r6, -1;
	.loc 3 210 5
	min.s32 	%r23, %r21, %r22;
	.loc 2 14 1
	sub.s32 	%r24, %r2, %r8;
	.loc 3 238 5
	max.s32 	%r25, %r24, %r20;
	.loc 2 14 1
	add.s32 	%r26, %r5, -1;
	.loc 3 210 5
	min.s32 	%r27, %r25, %r26;
	.loc 2 14 1
	mad.lo.s32 	%r28, %r23, %r5, %r27;
	sub.s32 	%r29, %r1, %r7;
	.loc 3 238 5
	max.s32 	%r30, %r29, %r20;
	.loc 2 14 1
	add.s32 	%r31, %r4, -1;
	.loc 3 210 5
	min.s32 	%r32, %r30, %r31;
	.loc 2 14 1
	mad.lo.s32 	%r33, %r28, %r4, %r32;
	mul.wide.s32 	%rd5, %r33, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mad.lo.s32 	%r34, %r3, %r5, %r2;
	mad.lo.s32 	%r35, %r34, %r4, %r1;
	mul.wide.s32 	%rd7, %r35, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;

BB0_2:
	.loc 2 16 2
	ret;
}


`
	shift_ptx_35 = `
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

.visible .entry shift(
	.param .u64 shift_param_0,
	.param .u64 shift_param_1,
	.param .u32 shift_param_2,
	.param .u32 shift_param_3,
	.param .u32 shift_param_4,
	.param .u32 shift_param_5,
	.param .u32 shift_param_6,
	.param .u32 shift_param_7
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<37>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [shift_param_0];
	ld.param.u64 	%rd4, [shift_param_1];
	ld.param.u32 	%r4, [shift_param_2];
	ld.param.u32 	%r5, [shift_param_3];
	ld.param.u32 	%r6, [shift_param_4];
	ld.param.u32 	%r7, [shift_param_5];
	ld.param.u32 	%r8, [shift_param_6];
	ld.param.u32 	%r9, [shift_param_7];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 3 9 1
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r12, %tid.x;
	mad.lo.s32 	%r1, %r10, %r11, %r12;
	.loc 3 10 1
	mov.u32 	%r13, %ntid.y;
	mov.u32 	%r14, %ctaid.y;
	mov.u32 	%r15, %tid.y;
	mad.lo.s32 	%r2, %r13, %r14, %r15;
	.loc 3 11 1
	mov.u32 	%r16, %ntid.z;
	mov.u32 	%r17, %ctaid.z;
	mov.u32 	%r18, %tid.z;
	mad.lo.s32 	%r3, %r16, %r17, %r18;
	.loc 3 13 1
	setp.lt.s32 	%p1, %r1, %r4;
	setp.lt.s32 	%p2, %r2, %r5;
	and.pred  	%p3, %p1, %p2;
	setp.lt.s32 	%p4, %r3, %r6;
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_2;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 14 1
	sub.s32 	%r19, %r3, %r9;
	mov.u32 	%r20, 0;
	.loc 4 238 5
	max.s32 	%r21, %r19, %r20;
	.loc 3 14 1
	add.s32 	%r22, %r6, -1;
	.loc 4 210 5
	min.s32 	%r23, %r21, %r22;
	.loc 3 14 1
	sub.s32 	%r24, %r2, %r8;
	.loc 4 238 5
	max.s32 	%r25, %r24, %r20;
	.loc 3 14 1
	add.s32 	%r26, %r5, -1;
	.loc 4 210 5
	min.s32 	%r27, %r25, %r26;
	.loc 3 14 1
	mad.lo.s32 	%r28, %r23, %r5, %r27;
	sub.s32 	%r29, %r1, %r7;
	.loc 4 238 5
	max.s32 	%r30, %r29, %r20;
	.loc 3 14 1
	add.s32 	%r31, %r4, -1;
	.loc 4 210 5
	min.s32 	%r32, %r30, %r31;
	.loc 3 14 1
	mad.lo.s32 	%r33, %r28, %r4, %r32;
	mul.wide.s32 	%rd5, %r33, 4;
	add.s64 	%rd6, %rd2, %rd5;
	ld.global.nc.f32 	%f1, [%rd6];
	mad.lo.s32 	%r34, %r3, %r5, %r2;
	mad.lo.s32 	%r35, %r34, %r4, %r1;
	mul.wide.s32 	%rd7, %r35, 4;
	add.s64 	%rd8, %rd1, %rd7;
	st.global.f32 	[%rd8], %f1;

BB2_2:
	.loc 3 16 2
	ret;
}


`
)
