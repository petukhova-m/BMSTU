import java.util.ArrayList;
import java.util.Scanner;
import java.util.Stack;

public class MaxComponent
{
    private static int countVer;
    private static int countEdge;


    private static class Vertex {
        ArrayList<Integer> edge = new ArrayList<>();
        int visit;
        int numComp;
    }

    public static void visitVertex(Vertex[] nodes, Vertex w, int comp) {
        Stack<Vertex> stack = new Stack<>();
        //for(Vertex w : nodes)
        //if (w.visit == 0){
        stack.push(w);
        while (stack.size() > 0) {
            Vertex v = stack.pop();
            for (Integer u : v.edge)
                if (nodes[u].visit == 0) {
                    nodes[u].visit = 1;
                    nodes[u].numComp = comp;
                    countVer++;
                    countEdge += nodes[u].edge.size();
                    stack.push(nodes[u]);
                }
        }
    }

    public static void main(String[] args) {
        long start = System.currentTimeMillis();
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        Vertex[] nodes = new Vertex[n];
        int m = in.nextInt();

        for (int i = 0; i < n; i++) {
            nodes[i] = new Vertex();
            nodes[i].visit = 0;
            nodes[i].numComp = -1;
        }
        for(int i = 0; i < m; i++){
            int u, v;
            u = in.nextInt();
            v = in.nextInt();
            nodes[u].edge.add(v);
            if (u != v){
                nodes[v].edge.add(u);
            }
        }



        int maxNodes = -1, maxEdge = -1, comp = 0, maxComp = 0;


        for(Vertex v : nodes){
            if(v.visit == 0){
                countEdge = 0;
                countVer = 0;
                visitVertex(nodes, v, comp);
                countEdge/=2;
                comp++;
            }
            if(countVer > maxNodes || (countVer == maxNodes && countEdge > maxEdge)){
                maxNodes = countVer;
                maxEdge = countEdge;
                maxComp = comp - 1;
            }
        }
        if(m == 0)
            nodes[0].numComp = 0;


        System.out.println("graph {");
        for(int i = 0; i < n; i++){
            System.out.printf("%d ", i);
            if(nodes[i].numComp == maxComp)
                System.out.print("[color = red]");
            System.out.print("\n");
        }
        //for(int i = 0; i < n; i++)
        //  nodes[i].visit = 0;
        //for(int i = 0; i < n; i++)
        for(int i = 0; i < n; i++)
            for(int j = 0; j <nodes[i].edge.size(); j++){
                int k = nodes[i].edge.get(j);
                if (i <= k){
                    System.out.printf("%d -- %d ", i, k);
                    if (nodes[i].numComp == maxComp)
                        System.out.print("[color = red]");
                    System.out.print("\n");
                }

            }
        System.out.println("}");
        //System.out.printf("%d", nodes[0].numComp);
        //System.out.println((double) (System.currentTimeMillis() - start));
    }
}