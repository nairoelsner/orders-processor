import { HttpException, Inject, Injectable } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';
import { AmqpConnection } from '@golevelup/nestjs-rabbitmq';

const PRODUCTS = [
  {id: 1, name: 'Product 1', price: 100, stock: 10},
  {id: 2, name: 'Product 2', price: 200, stock: 20},
  {id: 3, name: 'Product 3', price: 300, stock: 30},
  {id: 4, name: 'Product 4', price: 400, stock: 40},
  {id: 5, name: 'Product 5', price: 500, stock: 50},
];

const ORDERS: Order[] = [];

@Injectable()
export class OrdersService {
  constructor(private readonly amqpConnection: AmqpConnection) {}

  

  create(createOrderDto: CreateOrderDto) {
    if (createOrderDto.product < 1 || createOrderDto.product > 5) {
      throw new HttpException('Product not found', 404);
    }

    const newOrder = {
      id: ORDERS.length + 1,
      status: 'pending',
      ...createOrderDto,
    };
    
    ORDERS.push(newOrder);
    
    this.amqpConnection.publish('amq.topic', 'order.created', newOrder);
    
    return newOrder;
  }

  findAll() {
    return ORDERS;
  }

  findOne(id: number) {
    return ORDERS[id - 1];
  }

  update(id: number, updateOrderDto: UpdateOrderDto) {
    return `This action updates a #${id} order`;
  }

  remove(id: number) {
    return `This action removes a #${id} order`;
  }
}
