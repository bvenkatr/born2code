class Parent {
	constructor(name) {
		this.name = name;
	}
	
	spam() {
		console.log("Parent.spam %s", this.name);
	}
	
	grok() {
		console.log("Parent.grok");
		this.spam();
	}
}

class Child extends Parent {
	constructor(name) {
		super(name);
	}

	yow() {
		console.log("Child.yow");
	}	
}

module.exports = {
	Parent: Parent,
	Child: Child
}
