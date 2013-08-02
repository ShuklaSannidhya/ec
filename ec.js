var Subshell, ssn;

ssn = "spdfghijklmnoqrtuvwxyz";
Subshell = {
    n: 1,
    s: 's',
    maxE: 2,
    advance: function() {
        var nN, nS;
        
        if (this.s === 's') {
            nN = Math.ceil(this.n / 2) + 1;
            nS = ssn.charAt(nN - (this.n % 2) - 1);
        } else {
            nN = this.n + 1;
            nS = ssn.charAt(ssn.indexOf(this.s) - 1);
        }
        this.n = nN;
        this.s = nS;
        this.maxE = 2 * (2 * ssn.indexOf(nS) + 1);
    }
};

function getEC(aN) {
    var cSubshell, out;
    
    cSubshell = Object.create(Subshell);
    out = '';
    while (aN > 0) {
        out += cSubshell.n + cSubshell.s;
        if (aN - cSubshell.maxE > 0) {
            out += cSubshell.maxE;
        } else {
            out += aN;
        }
        aN -= cSubshell.maxE;
        cSubshell.advance();
        if (aN > 0) {
            out += ', ';
        }
    }
    return out;
}
