#include "textflag.h"

TEXT ·GetG(SB),NOSPLIT,$0-8
    MOVD g, res+0(FP)
    RET
