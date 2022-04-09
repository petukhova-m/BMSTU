class Vertex implements Comparable<Vertex>{
    int x, y;
    int dep;
    Vertex parent;

    public Vertex(int x, int y) {
        this.x = x;
        this.y = y;
        dep = 0;
        parent = this;
    }

    public int compareTo(Vertex obj){
        if (this.x > obj.x && this.y > obj.y)
            return 1;
        if (this.x == obj.x && this.y == obj.y)
            return 0;
        else return -1;
    }

    public Vertex find(){
        if(compareTo(this.parent) == 0)
            return this;
        else
            return this.parent.find();
    }
}