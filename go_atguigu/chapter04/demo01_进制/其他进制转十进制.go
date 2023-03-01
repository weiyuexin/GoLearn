package main

func main() {
	// 二进制转十进制
	// 规则：从最低位开始，将每个位上的数提取出来，乘以2的（位数-1）次方，然后求和
	// 1011 = 1 * 2^0 + 1 * 2^1 + 0 * 2^2 + 1 ^ 2^3 = 11

	// 八进制转十进制
	// 规则：从最低位开始，将每个位上的数提取出来，乘以8的（位数-1）次方，然后求和
	// 0123 = 3 * 8^0 + 2 * 8^1 + 1 * 8^2 = 83

	// 十六进制转十进制
	//从最低位开始，将每个位上的数提取出来，乘以16的（位数-1）次方，然后求和
	// 0x34A = 10 * 16^0 + 4 * 16^1 + 3 * 16^2 = 842

	// 练习
	// 二进制 110001100 转成 十进制
	// 110001100 = 1 * 2^2 + 1 * 2^3 + 1* 2^7 +1 * 2^8 = 4 + 8 +128 + 256 = 396
	// 八进制 02456 转成十进制
	// 02456 = 6 * 8^0 + 5 * 8^1 + 4 * 8^2 + 2 * 8^3 =6 + 40 + 256 + 1024 = 1326
	// 十六进制 0xA45 转成十进制
	// 0xA45 = 5 * 16^0 + 4 * 16^1 + 10 * 16^2 = 2629

}
