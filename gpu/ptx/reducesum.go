package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["reducesum"] = REDUCESUM }

const REDUCESUM = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_0000139b_00000000-9_reducesum.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/reducesum.cu"
.extern .shared .align 4 .b8 sdata[];

.visible .entry partial_sum(
	.param .u64 partial_sum_param_0,
	.param .u64 partial_sum_param_1
)
{
	.reg .pred 	%p<6>;
	.reg .s32 	%r<52>;
	.reg .s64 	%rd<17>;


	ld.param.u64 	%rd3, [partial_sum_param_0];
	ld.param.u64 	%rd4, [partial_sum_param_1];
	cvta.to.global.u64 	%rd1, %rd4;
	.loc 2 11 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r51, %ntid.x;
	mul.lo.s32 	%r6, %r1, %r51;
	shl.b32 	%r7, %r6, 1;
	.loc 2 10 1
	mov.u32 	%r3, %tid.x;
	.loc 2 11 1
	add.s32 	%r8, %r7, %r3;
	cvta.to.global.u64 	%rd5, %rd3;
	.loc 2 12 1
	mul.wide.u32 	%rd6, %r8, 4;
	add.s64 	%rd7, %rd5, %rd6;
	add.s32 	%r9, %r51, %r8;
	mul.wide.u32 	%rd8, %r9, 4;
	add.s64 	%rd9, %rd5, %rd8;
	ld.global.u32 	%r10, [%rd9];
	ld.global.u32 	%r12, [%rd7];
	add.s32 	%r14, %r10, %r12;
	mul.wide.u32 	%rd10, %r3, 4;
	mov.u64 	%rd11, sdata;
	add.s64 	%rd2, %rd11, %rd10;
	st.shared.u32 	[%rd2], %r14;
	.loc 2 13 1
	bar.sync 	0;
	.loc 2 15 1
	setp.lt.u32 	%p1, %r51, 66;
	@%p1 bra 	BB0_4;

BB0_1:
	.loc 2 15 1
	mov.u32 	%r4, %r51;
	shr.u32 	%r51, %r4, 1;
	.loc 2 16 1
	setp.ge.u32 	%p2, %r3, %r51;
	@%p2 bra 	BB0_3;

	.loc 2 17 1
	add.s32 	%r16, %r51, %r3;
	mul.wide.u32 	%rd12, %r16, 4;
	add.s64 	%rd14, %rd11, %rd12;
	ld.shared.u32 	%r17, [%rd2];
	ld.shared.u32 	%r19, [%rd14];
	add.s32 	%r21, %r17, %r19;
	st.shared.u32 	[%rd2], %r21;

BB0_3:
	.loc 2 19 1
	bar.sync 	0;
	.loc 2 15 1
	setp.gt.u32 	%p3, %r4, 131;
	@%p3 bra 	BB0_1;

BB0_4:
	.loc 2 22 1
	setp.gt.u32 	%p4, %r3, 31;
	@%p4 bra 	BB0_6;

	.loc 2 23 1
	ld.shared.u32 	%r23, [%rd2];
	ld.shared.u32 	%r25, [%rd2+128];
	add.s32 	%r27, %r23, %r25;
	st.shared.u32 	[%rd2], %r27;
	.loc 2 24 1
	ld.shared.u32 	%r29, [%rd2+64];
	add.s32 	%r31, %r27, %r29;
	st.shared.u32 	[%rd2], %r31;
	.loc 2 25 1
	ld.shared.u32 	%r33, [%rd2+32];
	add.s32 	%r35, %r31, %r33;
	st.shared.u32 	[%rd2], %r35;
	.loc 2 26 1
	ld.shared.u32 	%r37, [%rd2+16];
	add.s32 	%r39, %r35, %r37;
	st.shared.u32 	[%rd2], %r39;
	.loc 2 27 1
	ld.shared.u32 	%r41, [%rd2+8];
	add.s32 	%r43, %r39, %r41;
	st.shared.u32 	[%rd2], %r43;
	.loc 2 28 1
	ld.shared.u32 	%r45, [%rd2+4];
	add.s32 	%r47, %r43, %r45;
	st.shared.u32 	[%rd2], %r47;

BB0_6:
	.loc 2 32 1
	setp.ne.s32 	%p5, %r3, 0;
	@%p5 bra 	BB0_8;

	.loc 2 32 1
	ld.shared.u32 	%r49, [sdata];
	mul.wide.u32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd1, %rd15;
	st.global.u32 	[%rd16], %r49;

BB0_8:
	.loc 2 33 2
	ret;
}


`
