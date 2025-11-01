
var MyQueue = function() {
  this.stackPop = []
  this.stackPush = []

  this.transfer = function() {
    let pop = this.stackPush.pop()

    while (pop) {
      this.stackPop.push(pop)
      pop = this.stackPush.pop()
    }
  }
};

/**
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
  this.stackPush.push(x)
};

/**
 * @return {number}
 */
MyQueue.prototype.pop = function() {
  if (!this.stackPop.length) {
    this.transfer()
  }

  return this.stackPop.pop()
};

/**
 * @return {number}
 */
MyQueue.prototype.peek = function() {
  if (!this.stackPop.length) {
    this.transfer()
  }

  return this.stackPop.at(-1)
};

/**
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
  return !this.stackPop.length && !this.stackPush.length
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */