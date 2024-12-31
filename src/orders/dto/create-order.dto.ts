import { IsInt, IsString } from "class-validator";

export class CreateOrderDto {
    @IsString()
    customer: string;

    @IsInt()
    product: number;
}
