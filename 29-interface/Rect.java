interface IShape{
    float Area();
    float Perimeter();
   }
   
public class Rect implements IShape{
    float L,B;

    public float Area(){
        return this.L * this.B;
    }

    public float Perimeter(){
        return 2 * (this.L + this.B);
    }

}

