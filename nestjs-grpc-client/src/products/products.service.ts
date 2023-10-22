import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { Metadata } from '@grpc/grpc-js';
import { Client, ClientGrpc, Transport } from '@nestjs/microservices';
import { Observable, lastValueFrom } from 'rxjs';
import { join } from 'path';

type Product = {
  id: number,
  name: string,
  description: string,
  price: number
}

interface ProductGrpcClient {
  createProduct(
    data: {
      name: string,
      description: string,
      price: number
    }, 
    metadata?: Metadata,
  ): Observable<{ product: Product; }>;

  findProducts(
    data: {},
    metadata?: Metadata
  ): Observable<{ products: Product[] }>;
}

@Injectable()
export class ProductsService {
  @Client({
    transport: Transport.GRPC,
    options: {
      url: 'grpc:50051',
      package: 'github.com.codeedu.codepix',
      protoPath: [join(__dirname, 'protofiles', 'product.proto' )],
      loader: { keepCase: true },
    },
  })
  clientGrpc: ClientGrpc;

  private productGrpcClient: ProductGrpcClient;

  onModuleInit() {
    this.productGrpcClient = this.clientGrpc.getService('ProductService');
  }

  create(createProductDto: CreateProductDto) {
    return lastValueFrom(this.productGrpcClient.createProduct(createProductDto));
  }

  findAll() {
    try{
      return lastValueFrom(this.productGrpcClient.findProducts({}));
    } catch(e) {
        throw new HttpException(e, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

  // findOne(id: number) {
  //   return `This action returns a #${id} product`;
  // }

  // update(id: number, updateProductDto: UpdateProductDto) {
  //   return `This action updates a #${id} product`;
  // }

  // remove(id: number) {
  //   return `This action removes a #${id} product`;
  // }
}
