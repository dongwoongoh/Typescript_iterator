import createRangeIterable from "./01_createRangeIterable";
import { generator } from "./02_generator";

//01 iterator
const result_01 = createRangeIterable(1, 3 + 1);
while (true) {
  const { value, done } = result_01.next();
  if (done) break;
  console.log(value);
}
console.clear();

//02 generator
for (let i of generator()) console.log(i);
