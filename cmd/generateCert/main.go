/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/20
  @desc: //TODO
**/

package main

import (
	"flag"
	. "ginFlutterBolg/util"
)

func main() {
	fileType := flag.String("fileType", "cert", "Save the format")
	organization := flag.String("organization", "cilangOrganization", "cert organization")
	country := flag.String("country", "CN", "province Country")
	province := flag.String("province", "CN", "cert province")
	locality := flag.String("locality", "GuangZhou", "locality")
	organizationUnit := flag.String("organizationUnit", "cilang.buzz", "organizationUnit")
	commonName := flag.String("commonName", "cilang", "commonName")
	dnsDomain := flag.String("dnsDomain", "localhost", "dnsDomain")
	flag.Parse()
	GenerateCert(*fileType, *organization, *country, *province, *locality, *organizationUnit, *commonName, *dnsDomain)
	//GenerateCert("cert", "www.cilang.buzz", "CN",
	//	"GuangZhou", "GuangZhou", "cilang.buzz", "cilang", "localhost")
}
