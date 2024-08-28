// app.ts
import {Style} from './domain/styles/style'


const values: Map<string, string> = new Map();
values.set('first', 'fValue');
values.set('second', 'sValue');
values.set('third', 'tValue');

try {
    console.log(new Style(values).Fetch("fourth"));
} catch (e) {
    console.log(`error: ${e}`)
}

