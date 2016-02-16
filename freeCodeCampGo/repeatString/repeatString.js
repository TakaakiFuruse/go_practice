function repeat(str, num) {
    doRepeat = new Repeater(str, num)
    doRepeat.doRepeat()
}

function Repeater(str, num) {
    this.str = str
    this.num = num
}

Repeater.prototype.doRepeat = function() {
    var s = ""
    self = this
    for (var i = 1; i <= self.num; i++) {
        s += self.str
    };
    console.log(s)
    return s
}

repeat("*", 3)
repeat("abc", 3)
repeat("abc", 4)
repeat("abc", 1)
repeat("*", 8)
repeat("abc", -2)
