export namespace app {
	
	export enum Resolution {
	    RESOLUTION_2160 = "2160",
	    RESOLUTION_1440 = "1440",
	    RESOLUTION_1080 = "1080",
	    RESOLUTION_720 = "720",
	    RESOLUTION_480 = "480",
	    RESOLUTION_360 = "360",
	}
	export class Options {
	    url: string;
	    outputDir: string;
	    outputTempl: string;
	    targetResolution: Resolution;
	
	    static createFrom(source: any = {}) {
	        return new Options(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.outputDir = source["outputDir"];
	        this.outputTempl = source["outputTempl"];
	        this.targetResolution = source["targetResolution"];
	    }
	}

}

