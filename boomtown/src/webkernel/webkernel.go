package webkernel

import (

)

type Webkernel interface {
	
	Require(string,*RShell) string
	KernelInfo() string
	LinkRs(rs *RShell)
	UnLinkRs(key int)
			
}


