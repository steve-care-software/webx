import { stringify } from "querystring";

export class Style {
    private values: Map<string, string>;
    constructor(
        values: Map<string, string>, 
    ) {
        this.values = values;
    };

    public Fetch(name: string): string {
        const result: string | undefined =  this.values.get(name)
        if (result == undefined) {
            throw new Error(`the name ${name} could not be found`);
        }
        return result as string
    }
};