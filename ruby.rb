require 'net/http'

module ExampleModule
  class ParentClass
    def initialize
      puts("Parent Class")
    end
  end

  class ExampleClass < ParentClass

    def initialize(cv, iv)
      super()
      @@classvar = cv
      @instancevar = iv
    end

    def method1()
      self.class.method2()
      p_method()
    end

    def self.method2()
      if @@classvar <= 10
        puts "<= 10"
      end
    end

    private

    def p_method()
      x = {a: 5, b: 6}

      x.each do |k, v|
        puts "Key: #{k}, Value: #{v}"
      end
    end
  end
end

ec = ExampleModule::ExampleClass.new(1, 2)
ec.method1()
ExampleModule::ExampleClass.method2()
