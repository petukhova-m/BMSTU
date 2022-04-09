class Road implements Comparable<Road>{
    Vertex A;
    Vertex B;
    double len;

    public Road(Vertex A, Vertex B) {
        this.A = A;
        this.B = B;
        //len = 0;
    }

    @Override
    public int compareTo(Road obj) {
        if (this.len > obj.len)
            return 1;
        if (this.len == obj.len)
            return 0;
        else
            return -1;
    }
}