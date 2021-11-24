package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
)

func TestRsa(t *testing.T) {
	//生成密钥对，保存到文件
	//GenerateRSAKey(2048)
	//GenerateRSAKey(1024)

	message := []byte("hello world")
	//加密
	cipherText := RSA_Encrypt(message, "public.pem")
	fmt.Printf("加密后为：\n%s\n\n", string(cipherText))

	//解密
	plainText := RSA_Decrypt(cipherText, "private.pem")
	fmt.Printf("解密后为：\n%s\n\n", string(plainText))
}

//生成RSA私钥和公钥，保存到文件中
// bits 证书大小
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//将数据保存到内存
	pc := pem.EncodeToMemory(&privateBlock)
	str1 := string(pc)
	fmt.Printf("从内存取出生成的私钥\n\n%s\n", str1)

	//打开文件
	file, err := os.Open("private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	str2 := string(buf)
	fmt.Printf("比较从文件和内存取出生成的私钥数据是否一致 %t\n\n", str1 == str2)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "PUBLIC KEY", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)

}

//RSA公钥加密
func RSA_Encrypt(plainText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	//fmt.Printf("\npublicKey666rrrrf\n\n%x\n\n", string(buf))
	//
	//
	//strVal := "2d2d2d2d2d424547494e20525341205075626c6963204b65792d2d2d2d2d0a4d494942496a414e42676b71686b6947397730424151454641414f43415138414d49494243674b434151454176644c49673454534956547651717243377465420a3436415573687232536e6f2b7538563476424d69786573512b45577261744e6768624a6572467157526c2b2b474a7a6d75437366524b744f734b595a397054670a63566d383134356c54527834514b43326d5a5842675537776d34335769657a657a4646737465786d7a4e6f566c724b52374a4479555459774e386c74742b6f740a7578586377447a64634e79487a733031557231376b77396152386e68575359507a414273787774656a4a30486b4e6b586b564e50307930684f573542377543700a71356d786458714e68414b416a6a4152784e564c2b635275324d344c74415363756a657137774d494843754d49314c6a774b633464385461507a75306d3033630a593867505445574d4f53644a695a53327668745443614444303555774b75482f6d736d6c514f5462634d6679554d676f5a6e712f444d464c68567739536231690a68774944415141420a2d2d2d2d2d454e4420525341205075626c6963204b65792d2d2d2d2d"
	//buf2 := []byte(strVal)
	//
	//fmt.Printf("\npublicKey666rrrr99999f\n\n%x\n\n", buf2)

	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA私钥解密
func RSA_Decrypt(cipherText []byte, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	fmt.Printf("\n从文件中取出私钥进行解密 \n\n%s\n\n", string(buf))

	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}
