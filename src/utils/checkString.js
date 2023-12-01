export function isBetweenZeroAndTwenty(str) {
    // 首先使用正则表达式检查字符串是否为正整数
    if (/^\d+$/.test(str)) {
      // 将字符串转换为整数
      const number = parseInt(str, 10);
  
      // 检查数字是否在0到20之间
      if (number >= 0 && number <= 100) {
        return true; // 符合条件
      }
    }
    return false; // 不符合条件
  }