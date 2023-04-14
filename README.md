# Typescript iterator !

## 02.Generator

**"function\*" and "yield"**
src/02.generator.ts
Yield smliar return. but a some diff. basically, yield is wait return then complete a conditions.

```javascript
export function* generator() {
  console.log("Strat...🕐");
  let value = 1;
  while (value < 4) yield value++;
  console.log("Finished!");
}
```

src/index.ts

```javascript
import { generator } from "./02_generator";
for (let i of generator()) console.log(i);
```
