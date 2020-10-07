package main

func reverseString(s []byte)  {
	i,j :=0, len(s) -1
	for i <j{
		s[i],s[j] =s[j],s[i]
		//tmp:=s[i]
		//s[i]=s[j]
		//s[j]=tmp
		i++
		j--
	}
}

func main() {
	
}
