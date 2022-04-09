import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Scanner;

class Vertex{
    ArrayList<Integer> edge = new ArrayList<>();
    int parent;
    int component;
    int mark;
    int index;

    public Vertex(){
        this.component = -1;
        this.mark = -1;
        this.parent = -1;
    }
}

public class BridgeNum {
    private static int component;

    public static void dfs1(Vertex[] nodes, Queue queue){
        for(Vertex v : nodes){
            if(v.mark == -1) {
                component--;
                VisitVertex1(nodes, v, queue);
            }
            component = dfs2(nodes, queue);
        }
    }

    public static void VisitVertex1(Vertex[] nodes, Vertex v, Queue queue){
        v.mark = 0;
        queue.add(v);
        for(Integer u : v.edge)
            if(nodes[u].mark == -1){
                nodes[u].parent = v.index;
                VisitVertex1(nodes, nodes[u], queue);
            }
        v.mark = 1;
    }

    public static int dfs2(Vertex[] nodes, Queue queue){
        while(queue.size() > 0){
            Vertex v = (Vertex) queue.remove();
            if(v.component == -1){
                VisitVertex2(nodes, v);
                component++;
            }
        }
        return component;
    }

    public static void VisitVertex2(Vertex[] nodes, Vertex v){
        v.component = component;
        for(Integer u : v.edge)
            if(nodes[u].component == -1 && nodes[u].parent != v.index)
                VisitVertex2(nodes, nodes[u]);
    }

    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();
        int m = in.nextInt();
        Vertex nodes[] = new Vertex[n];
        for(int i = 0; i < n; i++) {
            nodes[i] = new Vertex();
            nodes[i].index = i;
        }
        for(int i = 0; i < m; i++){
            int u = in.nextInt();
            int v = in.nextInt();
            nodes[u].edge.add(v);
            nodes[v].edge.add(u);
        }

        Queue<Vertex> queue = new LinkedList<>();
        component = 1;
        dfs1(nodes, queue);
        System.out.println(component-1);

    }
}
