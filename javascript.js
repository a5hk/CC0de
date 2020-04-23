import { Dog } from "./modules/dog.js";
let myDog = new Dog("dog");
myDog.speak();

class Cat {
  constructor(name) {
    this.name = name;
  }

  speak() {
    console.log(`${this.name} makes a noise.`);
  }

  static eat() {
    console.log("I cannot eat");
  }
}

class Lion extends Cat {
  speak() {
    super.speak();
    console.log(`${this.name} roars.`);
  }
}

setTimeout(() => {
  let l = new Lion("Fuzzy");

  if (canISpeak()) {
    l.speak();
  }
}, 10);

function canISpeak() {
  let now = new Date();

  if (now.getHours() < 18 && now.getHours() > 10) {
    return true;
  }
  return false;
}
