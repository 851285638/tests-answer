#include "postgres.h"
#include "fmgr.h"
#include <math.h>

PG_FUNCTION_INFO_V1(simple_feistel_self_inverse);

Datum
simple_feistel_self_inverse(PG_FUNCTION_ARGS)
{
	int32 val = PG_GETARG_INT32(0);
	int32 l1 = (val >> 16) & 0xffff;
	int32 r1 = val & 0xffff;
	int32 l2, r2;
	int i;

	for (i = 0; i < 3; i++)
	{
		l2 = r1;
		/* round() is used to produce the same values as the
		   plpgsql implementation that does an SQL cast to INT */
		r2 = l1 ^ (int32)round((((1366*r1 + 150889) % 714025) / 714025.0) * 32767);
		l1 = l2;
		r1 = r2;
	}

	PG_RETURN_INT32((r1 << 16) + l1);
}