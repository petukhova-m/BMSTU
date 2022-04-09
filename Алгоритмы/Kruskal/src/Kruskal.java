import java.util.Arrays;
import java.util.Scanner;

public class Kruskal {

    static class Vertex implements Comparable<Vertex>{
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

    static class Road implements Comparable<Road>{
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

    public static void union(Vertex v, Vertex u){
        Vertex v1 = v.find();
        Vertex u1 = u.find();
        if(v1.dep > u1.dep)
            u1.parent = v1;
        else{
            v1.parent = u1;
            u1.dep++;
        }
    }

    public static void main(String[] args) {
        long start = System.currentTimeMillis();
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        Vertex[] nodes = new Vertex[n];
        for (int i = 0; i < n; i++) {
            int x = in.nextInt();
            int y = in.nextInt();
            nodes[i] = new Vertex(x, y);
        }

        int num = n * (n - 1) / 2;
        Road[] roads = new Road[num];
        int k = 0;
        for (int i = 0; i < n; i++)
            for (int j = i + 1; j < n; j++) {
                roads[k] = new Road(nodes[i], nodes[j]);
                roads[k].len = Math.sqrt(Math.pow(nodes[i].x - nodes[j].x, 2) + Math.pow(nodes[i].y - nodes[j].y, 2));
                k++;
            }
        Arrays.sort(roads);
            //for(int i = 0; i < num; i++)
                //System.out.println(roads[i].len);
         double summ = 0;
         int edges = 0;
         for(int i = 0; i < num && edges < n - 1; i++) {
             //System.out.println(roads[i].A.find().x + " " + roads[i].A.find().y);
             //System.out.println(roads[i].B.find().x + " " + roads[i].B.find().y);
             //System.out.println();
             if (roads[i].A.find().compareTo(roads[i].B.find()) != 0) {
                 summ += roads[i].len;
                 union(roads[i].A, roads[i].B);
                 edges++;
             }
         }
          System.out.printf("%.2f\n", summ);
         System.out.println((double)(System.currentTimeMillis() - start));
    }
}
