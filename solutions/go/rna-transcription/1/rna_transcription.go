package strand

func ToRNA(dna string) string {
    temp, rna := "", ""
    
    if dna == "" {
        return temp
    }
	
    for _, char := range(dna) {
        
        switch char {
            case 'G' : {
                temp = "C"
            }
        	case 'C' : {
                temp = "G"
            }
            case 'T' : {
                temp = "A"
            }
            default: {
                temp = "U"
            }
        }

        rna += temp
    }
    
	
	return rna
}
