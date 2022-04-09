import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Scanner;

public class EqDist {
    private static class Vertex{
        ArrayList<Integer> edge = new ArrayList<>();
        ArrayList<Integer> dist = new ArrayList<>();
    }

    private static void VisitVertex(Vertex[] nodes, int opora, int i, int[] visit){
        nodes[i].dist.set(opora, 0);
        for (int j = 0; j < nodes.length; j++)
            visit[j] = 0;
        Queue q = new LinkedList<Integer>();
        q.add(opora);
        while(q.size() > 0){
            int h = (int) q.remove();
            for(int j = 0; j < nodes[h].edge.size(); j++){
                if (visit[nodes[h].edge.get(j)] == 0){
                    visit[nodes[h].edge.get(j)] = 1;
                    q.add(nodes[h].edge.get(j));
                    nodes[i].dist.set(nodes[h].edge.get(j), nodes[i].dist.get(h) + 1);
                }
            }
        }
        //nodes[i].dist.set(opora, 0);
    }

    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int m = in.nextInt();
        Vertex[] nodes = new Vertex[n];
        int[] visit = new int[n];
        for(int i = 0; i < n; i++) {
            nodes[i] = new Vertex();
        }

        for(int i = 0; i < m; i++){
            int u = in.nextInt();
            int v = in.nextInt();
            nodes[u].edge.add(v);
            if(u!= v)
                nodes[v].edge.add(u);
        }
        int k = in.nextInt();
        int[] opora = new int[k];
        for(int i = 0; i < k; i++){
            opora[i] = in.nextInt();
            for(int j = 0; j < n; j++) {
                nodes[i].dist.add(-1);
            }
            VisitVertex(nodes, opora[i], i, visit);
            nodes[i].dist.set(opora[i], 0);
        }
        for(int i = 0; i < k; i++){
            for(int j = 0; j < n; j++)
                System.out.printf("%d ", nodes[i].dist.get(j));
            System.out.printf("\n");
        }
        int s = 0;
        for(int i = 0; i < n; i++){
            int mark = 1;
            for(int j = 0; j < k -1; j++)
                if(nodes[j].dist.get(i) == -1 || nodes[j].dist.get(i) != nodes[j+1].dist.get(i)){
                    mark = 0;
                    //break;
                }
            if (mark == 1) {
                s = 1;
                System.out.printf("%d ", i);
            }
        }
        if ( s == 0)
            System.out.printf("-");
    }
}
